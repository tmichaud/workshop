package com.tem.turnquist.learningspringboot.learningspringbootvideo;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Bean;
import org.springframework.core.io.Resource;
import org.springframework.core.io.ResourceLoader;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.util.FileCopyUtils;
import org.springframework.util.FileSystemUtils;
import org.springframework.web.multipart.MultipartFile;

// Disk store and DB
@Service
public class ImageService
{
	private static String UPLOAD_ROOT = "upload-dir";

	private final ImageRepository repository;
	private final ResourceLoader resourceLoader;

	@Autowired
	public ImageService(ImageRepository repository, ResourceLoader resourceLoader)
	{
		this.repository = repository;
		this.resourceLoader = resourceLoader;
	}

	public Page<Image> findPage(Pageable pageable)
	{
		return repository.findAll(pageable);
	}

	public Resource findOneImage(String filename)
	{
		return resourceLoader.getResource("file:" + UPLOAD_ROOT + "/" + filename);
	}

	public void createImage(MultipartFile file) throws IOException
	{
		if (!file.isEmpty())
		{
			Files.copy(file.getInputStream(), Paths.get(UPLOAD_ROOT, file.getOriginalFilename()));
			repository.save(new Image(file.getOriginalFilename()));
		}
	}

	public void deleteImage(String filename) throws IOException
	{
		final Image byName = repository.findByName(filename);
		repository.delete(byName);
		Files.deleteIfExists(Paths.get(UPLOAD_ROOT, filename));
	}

	// Runs automatically to initialize data
	@Bean
	// @Profile("dev") // Only run this in dev
	CommandLineRunner setUp(ImageRepository repository)
	{
		return (args) -> {
			FileSystemUtils.deleteRecursively(new File(UPLOAD_ROOT));

			Files.createDirectory(Paths.get(UPLOAD_ROOT));

			FileCopyUtils.copy("Test file", new FileWriter(UPLOAD_ROOT + "/test"));
			repository.save(new Image("test"));

			FileCopyUtils.copy("Test file2", new FileWriter(UPLOAD_ROOT + "/test2"));
			repository.save(new Image("test2"));

			FileCopyUtils.copy("Test file3", new FileWriter(UPLOAD_ROOT + "/test3"));
			repository.save(new Image("test3"));
		};
	}
}
